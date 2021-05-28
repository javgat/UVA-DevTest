import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Mensaje } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-mensaje',
  templateUrl: './mensaje.component.html',
  styleUrls: ['./mensaje.component.css']
})
export class MensajeComponent implements OnInit {

  private mensaje: Mensaje
  private messageSubscription: Subscription

  constructor(private data: DataService) {
    this.mensaje = new Mensaje()
    this.messageSubscription = this.data.mensajeActual.subscribe(
      valor => this.mensaje = valor
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
}
