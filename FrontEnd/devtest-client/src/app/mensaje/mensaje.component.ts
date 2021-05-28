import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-mensaje',
  templateUrl: './mensaje.component.html',
  styleUrls: ['./mensaje.component.css']
})
export class MensajeComponent implements OnInit {

  private mensaje: Mensaje
  private messageSubscription: Subscription
  private mostrar: boolean
  private timerId?: number
  constructor(private data: DataService) {
    this.mostrar = true
    this.mensaje = new Mensaje()
    this.messageSubscription = this.data.mensajeActual.subscribe(
      valor => {
        this.mensaje = valor
        this.mostrar = true
        this.programarOcultacion()
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.messageSubscription.unsubscribe()
  }
  
  getMensaje(): Mensaje {
    return this.mensaje
  }

  showMensaje(): boolean{
    return this.getMensaje().mostrar && this.mostrar
  }

  ocultarMensaje(){
    this.mostrar = false
  }

  ocultarMensajeOther(that: MensajeComponent){
    that.ocultarMensaje()
  }

  programarOcultacion(){
    if(this.timerId!=undefined){
      clearTimeout(this.timerId);
      this.timerId = undefined
    }
    console.log(this.getMensaje().trueType)
    if(this.getMensaje().trueType == Tipo.SUCCESS){
      console.log("en 3 secs adios")
      this.timerId = setTimeout(this.ocultarMensajeOther, 2500, this)
    }
  }
}
