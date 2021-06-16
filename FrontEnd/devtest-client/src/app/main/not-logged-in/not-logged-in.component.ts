import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { MainComponent } from '../main.component';

import { ConfigurationService, CustomizedView, UserService } from '@javgat/devtest-api'
import { ActivatedRoute, Router } from '@angular/router';
import { VistaPersonalizada } from 'src/app/shared/app.model';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-not-logged-in',
  templateUrl: './not-logged-in.component.html',
  styleUrls: ['./not-logged-in.component.css']
})
export class NotLoggedInComponent extends MainComponent implements OnInit {

  vistaPers: CustomizedView
  messSub: Subscription
  constructor(session: SessionService, router: Router, datos: DataService, userService: UserService, private route: ActivatedRoute, private configS?: ConfigurationService) {
    super(session, router, datos, userService, route);
    this.messSub = datos.mensajeActual.subscribe(
      valor => {
        if (valor.mostrar == true) {
          this.borrarMensaje()
        }
      }
    )
    this.vistaPers = new VistaPersonalizada()
    this.getVistaPersonalizada(true)
  }

  ngOnInit(): void {
    this.route.fragment.subscribe(
      (fragments) => {
        if (fragments != undefined) {
          var elem = document.getElementById(fragments)
          if (elem != undefined) {
            this.router.navigate(['/'], { fragment: fragments })
          }else{
            setTimeout(() => {
              elem = document.getElementById(fragments)
              this.router.navigate(['/'], { fragment: fragments })
            }, 500);
          }
        }
      }
    );
  }

  ngOnDestroy(): void {
    this.messSub.unsubscribe()
    super.ngOnDestroy()
  }

  hasPermissions() {
    return true
  }

  getVistaPersonalizada(primera: boolean) {
    if (this.configS == undefined) return
    this.configS.getCView("noRegistrado").subscribe(
      resp => {
        this.vistaPers = new VistaPersonalizada(resp)
      },
      err => this.handleErrRelog(err, "obtener datos personalizados de la vista segun el rol", primera, this.getVistaPersonalizada, this)
    )
  }

}
