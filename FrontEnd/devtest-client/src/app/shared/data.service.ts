import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { Mensaje, Tipo } from './app.model'

// DataService mantiene datos globales que se mantienen entre 
// elementos de la navegacion

@Injectable({
  providedIn: 'root'
})
export class DataService {
  // Usamos mensajes para mostrar el resultado de la operacion
  private mensaje = new BehaviorSubject<Mensaje>(new Mensaje());
  mensajeActual = this.mensaje.asObservable();

  constructor() { }

  // Cambia el mensaje a mostrar por el pasado por parametro
  cambiarMensaje(mensaje: Mensaje){
    this.mensaje.next(mensaje);
  }

  // Modifica el texto del mensaje actual
  cambiarTextoMensaje(texto: string) {
    let oldMsn = this.mensaje.value;
    this.mensaje.next(new Mensaje(texto, oldMsn.type, oldMsn.mostrar));
  }

  // Modifica el boolean de visibilidad del mensaje actual
  cambiarMostrarMensaje(mostrar: boolean) {
    let oldMsn = this.mensaje.value;
    this.mensaje.next(new Mensaje(oldMsn.texto, oldMsn.type, mostrar));
  }

  // Modifica el tipo del mensaje actual
  cambiarTypeMensaje(type : Tipo){
    let oldMsn = this.mensaje.value;
    this.mensaje.next(new Mensaje(oldMsn.texto, type, oldMsn.mostrar));
  }

  // Borra el mensaje.
  borrarMensaje(){
    this.cambiarMensaje(new Mensaje());
  }
}