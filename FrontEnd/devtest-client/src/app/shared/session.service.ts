import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Session } from './app.model';

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  private session = new BehaviorSubject<Session>(new Session(false))
  sessionActual = this.session.asObservable()

  constructor() { }

  cambiarSession(session:Session){
    this.session.next(session)
  }

  borrarSession(){
    this.session.next(new Session(false))
  }
}
