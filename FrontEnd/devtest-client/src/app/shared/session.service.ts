import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Session } from './app.model';

// SessionService aporta información global sobre la sesión actual,
// el jwt que tiene que transmitir, etc.

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  private session = new BehaviorSubject<Session>(new Session(false))
  sessionActual = this.session.asObservable()

  constructor() { }

  // Actualiza la sesión a la pasada por parametro.
  cambiarSession(session:Session){
    this.session.next(session)
  }

  // Desautentica al usuario. Elimina la sesión.
  borrarSession(){
    this.session.next(new Session(false))
  }
}
