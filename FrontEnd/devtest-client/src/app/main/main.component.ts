import { Component, OnInit } from '@angular/core';
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

  constructor(private datos: DataService, private session: SessionService) {
    this.sessionActual = new SessionUser(false)
  }

  ngOnInit(): void {
    this.session.sessionActual.subscribe(
      valor => this.sessionActual = valor
    )
  }

  logout(){
    this.session.borrarSession()
  }

}
