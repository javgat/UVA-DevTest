import { HttpHandler, HttpHeaderResponse, HttpInterceptor, HttpProgressEvent, HttpRequest, HttpResponse, HttpSentEvent, HttpUserEvent } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { Mensaje, SessionLogin, Tipo } from './app.model';
import { DataService } from './data.service';
import { SessionService } from './session.service';

@Injectable({
  providedIn: 'root'
})
export class CustomHttpInterceptorService implements HttpInterceptor {

  sessionLogin: SessionLogin
  sessionSubscription: Subscription
  constructor(private session: SessionService, private data: DataService) {
    this.sessionLogin = new SessionLogin(false)
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionLogin = valor
      }
    )
  }

  intercept(req: HttpRequest<any>, next: HttpHandler):
    Observable<HttpSentEvent | HttpHeaderResponse | HttpProgressEvent | HttpResponse<any> | HttpUserEvent<any>> {
    switch(req.method){
      case "PUT":
      case "DELETE":
      case "POST": {
        this.data.cambiarMensaje(new Mensaje("Enviando datos...", Tipo.SENDING, true))
        break;
      }
      case "GET":{
        this.data.borrarMensajeIfLoading()
        break;
      }
      default: {
        break;
      }
    }
    let notLoggedIn = 'true'
    if (this.sessionLogin.isLoggedIn()) {
      notLoggedIn = 'false'
    }
    const nextReq = req.clone({
      headers: req.headers.set('Cache-Control', 'no-cache')
        .set('Pragma', 'no-cache')
        .set('Expires', 'Sat, 01 Jan 2000 00:00:00 GMT')
        .set('If-Modified-Since', '0')
        .set('NotLoggedIn', notLoggedIn)
    });

    return next.handle(nextReq);
  }
}