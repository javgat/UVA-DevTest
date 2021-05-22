import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
import { DataService } from '../../shared/data.service';
import { SessionService } from '../../shared/session.service';
import { NavBarComponent } from '../nav-bar.component';

@Component({
  selector: 'app-nav-bar-not-logged',
  templateUrl: './nav-bar-not-logged.component.html',
  styleUrls: ['../nav-bar.component.css']
})
export class NavBarNotLoggedComponent extends NavBarComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS : UserService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }

}
