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

  //res.cookie("SESSIONID", jwtBearerToken, {httpOnly:true, secure:true});
  
  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session:SessionUser){
    this.session.next(session)
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession(){
    this.session.next(new SessionUser(false))
  }
}
