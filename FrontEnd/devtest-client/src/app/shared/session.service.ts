import { Injectable } from '@angular/core';
import { User } from '@javgat/devtest-api';
import { BehaviorSubject } from 'rxjs';
import { SessionLogin, SessionUser } from './app.model';

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

  constructor() { }
  
  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session:SessionLogin){
    localStorage.setItem('logged', String(session.logged))
    localStorage.setItem('userid', String(session.userid))
    this.session.next(session)
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession(){
    this.cambiarSession(new SessionLogin(false))
  }

  logout(){
    this.borrarSession()
    this.borrarUser()
  }

  checkStorageSession(){
    var logged = localStorage.getItem('logged')
    var userid = localStorage.getItem('userid')
    var loggedBool
    if(logged == null || userid == null){
      userid = "null"
      loggedBool = false
    }else{
      loggedBool = ("true"==logged)
    }
    this.cambiarSession(new SessionLogin(loggedBool, userid))
  }

  cambiarUser(user:SessionUser){
    this.user.next(user)
  }

  borrarUser(){
    this.cambiarUser(new SessionUser())
  }
}
