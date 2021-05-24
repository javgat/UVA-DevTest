import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConfigurationService, CustomizedView, TeamService, UserService } from '@javgat/devtest-api';
import { Mensaje, Tipo } from 'src/app/shared/app.model';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-customized-views',
  templateUrl: './admin-customized-views.component.html',
  styleUrls: ['./admin-customized-views.component.css']
})
export class AdminCustomizedViewsComponent extends AdminComponent implements OnInit {

  cvistas: CustomizedView[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService, private configS: ConfigurationService) {
    super(session, router, data, userS, teamS)
    this.cvistas = []
    this.getCViews(true)
  }

  ngOnInit(): void {
  }

  getCViews(primera: boolean){
    this.configS.getCViews().subscribe(
      resp => this.cvistas = resp,
      err => this.handleErrRelog(err, "obtener vistas personalizadas", primera, this.getCViews, this)
    )
  }

  submitPutMensajes(){
    this.putMensajes(true)
  }

  putMensajes(primera: boolean){
    for(let cvista of this.cvistas){
      this.configS.putCView(cvista.rolBase, cvista).subscribe(
        resp => {
          this.getCViews(true)
          this.cambiarMensaje(new Mensaje("Vistas personalizadas actualizadas con éxito", Tipo.SUCCESS, true))
        },
        err => this.handleErrRelog(err, "actualizar vista personalizada", primera, this.putMensajes, this)
      )
    }
  }

}
