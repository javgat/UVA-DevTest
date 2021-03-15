import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { SessionUser } from '../shared/app.model';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  sessionUser : SessionUser
  private sessionUserSubscription : Subscription
  
  constructor(private session: SessionService) {
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor => this.sessionUser = valor
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.sessionUserSubscription.unsubscribe();
  }

}
