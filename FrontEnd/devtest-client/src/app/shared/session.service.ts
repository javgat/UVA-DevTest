import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { SessionUser } from './app.model';

// SessionService aporta información global sobre la sesión actual,
// el jwt que tiene que transmitir, etc.

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  private session = new BehaviorSubject<SessionUser>(new SessionUser(false))
  sessionActual = this.session.asObservable()

  constructor() { }
  
  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session:SessionUser){
    this.session.next(session)
    localStorage.setItem('logged', String(session.logged))
    localStorage.setItem('userid', String(session.userid))
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession(){
    this.cambiarSession(new SessionUser(false))
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
    this.cambiarSession(new SessionUser(loggedBool, userid))
  }
}
