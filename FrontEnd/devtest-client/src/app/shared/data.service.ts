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
  cambiarMensaje(mensaje: Mensaje) {
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
  cambiarTypeMensaje(type: Tipo) {
    let oldMsn = this.mensaje.value;
    this.mensaje.next(new Mensaje(oldMsn.texto, type, oldMsn.mostrar));
  }

  // Borra el mensaje.
  borrarMensaje() {
    this.cambiarMensaje(new Mensaje());
  }

  // Borra el mensaje si esta cargando.
  borrarMensajeIfLoading() {
    if(this.mensaje.value.isLoading())
      this.borrarMensaje()
  }

  handleShowErr(err: any, action: string) {
    let msg: string
    if (err.error != undefined && err.error.message != undefined)
      msg = err.error.message
    else if (err.status >= 500)
      msg = "Error al conectar con el servidor"
    else if (err.status == 410 || err.status == 404)
      msg = "El recurso no existe o no está disponible"
    else if (err.status == 403)
      msg = "No tienes permiso para acceder a esta operación"
    else if (err.status == 409)
      msg = "Has intentado crear un recurso con un identificador ya utilizado"
    else
      msg = ""
    msg = "Error en " + action + ": " + err.status + " " + msg
    this.cambiarMensaje(new Mensaje(msg, Tipo.ERROR, true))
    console.log(msg)
  }
}