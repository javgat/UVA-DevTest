import { Injectable } from '@angular/core';
import { AuthService, TipoRol, TiporolService } from '@javgat/devtest-api';
import { BehaviorSubject } from 'rxjs';
import { SessionLogin, SessionUser } from './app.model';
import { DataService } from './data.service';

// SessionService aporta información global sobre la sesión actual,
// el jwt que tiene que transmitir, etc.

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  private session = new BehaviorSubject<SessionLogin>(new SessionLogin(false))
  sessionLogin = this.session.asObservable()

  private user = new BehaviorSubject<SessionUser>(new SessionUser())
  sessionUser = this.user.asObservable()

  private tipoRoles = new BehaviorSubject<TipoRol[]>([])
  sessionTipoRoles = this.tipoRoles.asObservable()

  constructor(private auth: AuthService, private data: DataService, private trS: TiporolService) { }
  
  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session:SessionLogin){
    localStorage.setItem('logged', String(session.isLoggedIn()))
    localStorage.setItem('username', String(session.getUserUsername()))
    this.session.next(session)
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession(){
    this.cambiarSession(new SessionLogin(false))
    this.auth.logout().subscribe(
      _ => console.log("Sesion cerrada con exito"),
      _ => console.log("Error al cerrar sesion")
    )
  }

  logout(){
    this.borrarSession()
    this.borrarUser()
  }

  checkStorageSession(){
    var logged = localStorage.getItem('logged')
    var username = localStorage.getItem('username')
    var loggedBool
    if(logged == null || username == null){
      username = ""
      loggedBool = false
    }else{
      loggedBool = ("true"==logged)
    }
    this.cambiarSession(new SessionLogin(loggedBool, username))
    this.updateTipoRoles(true)
  }

  cambiarUser(user:SessionUser){
    this.user.next(user)
  }

  borrarUser(){
    this.cambiarUser(new SessionUser())
  }
  
  cambiarSessionTipoRoles(ntrs: TipoRol[]){
    this.tipoRoles.next(ntrs)
  }

  updateTipoRoles(primera: boolean){
    this.trS.getTipoRoles().subscribe(
      resp=>{
        this.cambiarSessionTipoRoles(resp)
      },
      err=> this.handleErrRelog(err, "obtener tipo de roles", primera, this.updateTipoRoles, this)
    )
  }

  handleErrRelog<T>(err: any, action: string, primera: boolean, callbackFn: (this: T, prim: boolean) => void, that: T): void{
    if (err.status==401){
      if(!primera){
        this.data.handleShowErr(err, "alargar sesión de usuario")
        this.logout()
      }else{
        this.auth.relogin().subscribe(
          resp =>{
            return callbackFn.call(that, false)
          },
          err => {
            this.data.handleShowErr(err, "alargar sesión de usuario")
            this.logout()
          }
        )
      }
    }else{
      this.data.handleShowErr(err, action)
    }
  }

}
