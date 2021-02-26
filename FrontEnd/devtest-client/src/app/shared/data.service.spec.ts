import { TestBed } from '@angular/core/testing';
import { Mensaje, Tipo } from './app.model';

import { DataService } from './data.service';

describe('DataService', () => {
  let service: DataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DataService);    
    service.cambiarMensaje(new Mensaje())
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('#cambiarMensaje should change the messsage', () => {
    const mens = new Mensaje("hola", Tipo.ERROR, true)
    service.cambiarMensaje(mens)
    service.mensajeActual.subscribe(
      valor =>{
        expect(valor.mostrar).toBe(mens.mostrar)
        expect(valor.texto).toBe(mens.texto)
        expect(valor.type).toBe(mens.type)
      }
    )
  })
});
