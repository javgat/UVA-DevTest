import { Component, OnInit } from '@angular/core';
import { UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { SessionLogin } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  
  sessionLogin : SessionLogin
  private sessionSubscription : Subscription

  constructor(private datos: DataService, protected session: SessionService, protected userService : UserService) {
    this.sessionLogin = new SessionLogin(false)
    this.session.checkStorageSession()
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionLogin = valor
        //console.log(valor)//hay varias instancias de main component a la vez, porque se llama a si mismo
      }
    )
  }

  ngOnDestroy(): void {
    this.sessionSubscription.unsubscribe();
  }

  ngOnInit(): void {
    
  }

}
