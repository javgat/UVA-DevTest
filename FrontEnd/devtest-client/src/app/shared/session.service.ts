import { Injectable } from '@angular/core';
import { AuthService, TipoRol, TiporolService, UserService } from '@javgat/devtest-api';
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

  private localStorageLoggedKey = 'logged'
  private localStorageUsernameKey = 'username'
  private localStorageTipoRolesKey = 'tipoRoles'

  private updatingTipoRoles: boolean
  private updatingUser: boolean

  constructor(private auth: AuthService, private data: DataService, private trS: TiporolService, private userS: UserService) {
    this.updatingTipoRoles = false
    this.updatingUser = false
  }

  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session: SessionLogin) {
    localStorage.setItem(this.localStorageLoggedKey, String(session.isLoggedIn()))
    localStorage.setItem(this.localStorageUsernameKey, String(session.getUserUsername()))
    this.session.next(session)
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession() {
    this.cambiarSession(new SessionLogin(false))
    this.auth.logout().subscribe(
      _ => console.log("Sesion cerrada con exito"),
      _ => console.log("Error al cerrar sesion")
    )
  }

  logout() {
    this.borrarSession()
    this.borrarUser()
  }

  checkStorageSession() {
    var logged = localStorage.getItem(this.localStorageLoggedKey)
    var username = localStorage.getItem(this.localStorageUsernameKey)
    var loggedBool
    if (logged == null || username == null) {
      username = ""
      loggedBool = false
    } else {
      loggedBool = ("true" == logged)
    }
    this.cambiarSession(new SessionLogin(loggedBool, username))
    var tipoRolesStorage = this.getWithExpiry(this.localStorageTipoRolesKey)
    if (tipoRolesStorage == null) {
      this.updateTipoRoles(true)
    } else {
      this.cambiarSessionTipoRoles(tipoRolesStorage as TipoRol[])
    }
  }

  // ttl = milliseconds
  setWithExpiry(key: string, value: any, ttl: number) {
    const now = new Date()
    const item = {
      value: value,
      expiry: now.getTime() + ttl,
    }
    localStorage.setItem(key, JSON.stringify(item))
  }

  // if the item doesn't exist, return null
  getWithExpiry(key: string) {
    const itemStr = localStorage.getItem(key)
    if (!itemStr) {
      return null
    }
    const item = JSON.parse(itemStr)
    const now = new Date()
    if (now.getTime() > item.expiry) {
      localStorage.removeItem(key)
      return null
    }
    return item.value
  }

  cambiarUser(user: SessionUser) {
    this.user.next(user)
  }

  borrarUser() {
    this.cambiarUser(new SessionUser())
  }

  cambiarSessionTipoRoles(ntrs: TipoRol[]) {
    this.tipoRoles.next(ntrs)
  }

  updateTipoRoles(primera: boolean) {
    if (this.updatingTipoRoles) return
    this.updatingTipoRoles = true
    this.trS.getTipoRoles().subscribe(
      resp => {
        this.cambiarSessionTipoRoles(resp)
        this.setWithExpiry(this.localStorageTipoRolesKey, resp, 5 * 60000) // lo guarda 5 minutos
      },
      err => this.handleErrRelog(err, "obtener tipo de roles", primera, this.updateTipoRoles, this),
      () => this.updatingTipoRoles = false
    )
  }

  updateUser(primera: boolean) {
    if (this.updatingUser) return
    this.updatingUser = true
    this.userS.getUser(this.session.value.getUserUsername() as string).subscribe(
      resp => {
        this.cambiarUser(new SessionUser(resp.username, resp.email, resp.fullname, resp.rol, resp.tiporol))
      },
      err => {
        this.handleErrRelog(err, "obtencion del usuario", primera, this.updateUser, this)
      },
      () => this.updatingUser = false
    )
  }

  handleErrRelog<T>(err: any, action: string, primera: boolean, callbackFn: (this: T, prim: boolean) => void, that: T): void {
    if (err.status == 401) {
      if (!primera) {
        this.data.handleShowErr(err, "alargar sesión de usuario")
        this.logout()
      } else {
        this.auth.relogin().subscribe(
          resp => {
            return callbackFn.call(that, false)
          },
          err => {
            this.data.handleShowErr(err, "alargar sesión de usuario")
            this.logout()
          }
        )
      }
    } else {
      this.data.handleShowErr(err, action)
    }
  }

}
