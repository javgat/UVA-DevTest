import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent extends LoggedInController implements OnInit {

  constructor(session: SessionService, router: Router, datos: DataService, userS : UserService) {
    super(session, router, datos, userS)
  }

  ngOnDestroy(): void {
    this.borrarMensaje()
    super.onDestroy()
  }

  ngOnInit(): void {
    
  }

}
