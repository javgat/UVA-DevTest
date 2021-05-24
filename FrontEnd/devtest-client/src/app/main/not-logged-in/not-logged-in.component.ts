import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { MainComponent } from '../main.component';

import { ConfigurationService, CustomizedView, UserService} from '@javgat/devtest-api'
import { Router } from '@angular/router';
import { VistaPersonalizada } from 'src/app/shared/app.model';

@Component({
  selector: 'app-not-logged-in',
  templateUrl: './not-logged-in.component.html',
  styleUrls: ['./not-logged-in.component.css']
})
export class NotLoggedInComponent extends MainComponent implements OnInit {

  vistaPers: CustomizedView
  constructor(session: SessionService, router: Router, datos: DataService, userService : UserService, private configS?: ConfigurationService){
    super(session, router, datos, userService);
    this.vistaPers = new VistaPersonalizada()
    if(!this.getSessionUser().isEmpty()){
      this.getVistaPersonalizada(true)
    }
  }

  ngOnInit(): void {
  }

  doHasUserAction(){
    if(this.configS != undefined){
      this.getVistaPersonalizada(true)
    }
  }

  hasPermissions(){
    return true
  }

  getVistaPersonalizada(primera: boolean){
    if(this.configS==undefined) return
    this.configS.getCView("noRegistrado").subscribe(
      resp => {
        this.vistaPers = new VistaPersonalizada(resp)
      },
      err => this.handleErrRelog(err, "obtener datos personalizados de la vista segun el rol", primera, this.getVistaPersonalizada, this)
    )
  }

}
