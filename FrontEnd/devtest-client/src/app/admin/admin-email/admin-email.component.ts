import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConfigurationService, EmailConfiguration, TeamService, UserService } from '@javgat/devtest-api';
import { ConfiguracionCorreo } from 'src/app/shared/app.model';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-email',
  templateUrl: './admin-email.component.html',
  styleUrls: ['./admin-email.component.css']
})
export class AdminEmailComponent extends AdminComponent implements OnInit {

  emailConfig: EmailConfiguration
  emailConfigEdit: EmailConfiguration
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService, private configS: ConfigurationService) {
    super(session, router, data, userS, teamS)
    this.emailConfig = new ConfiguracionCorreo()
    this.emailConfigEdit = new ConfiguracionCorreo()
    this.getConfiguracionCorreo(true)
  }

  ngOnInit(): void {
  }

  getConfiguracionCorreo(primera: boolean){
    this.configS.getEmailConfiguration().subscribe(
      resp => {
        this.emailConfig = new ConfiguracionCorreo(resp)
        this.emailConfigEdit = new ConfiguracionCorreo(resp)
      },
      err => this.handleErrRelog(err, "obtener configuracion de correo", primera, this.getConfiguracionCorreo, this)
    )
  }

}
