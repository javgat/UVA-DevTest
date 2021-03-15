import { Component, OnInit } from '@angular/core';
import { UserService } from '@javgat/devtest-api';
import { SessionUser } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  
  sessionActual : SessionUser

  constructor(private datos: DataService, private session: SessionService, protected userService : UserService) {
    this.sessionActual = new SessionUser(false)
    this.userService.configuration.withCredentials = true
    this.session.checkStorageSession()
    this.session.sessionActual.subscribe(
      valor => this.sessionActual = valor
    )
  }

  ngOnInit(): void {
    
  }

  logout(){
    this.session.borrarSession()
  }

}
