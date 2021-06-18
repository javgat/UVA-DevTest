import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, Test, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ProfileComponent } from '../profile.component';

@Component({
  selector: 'app-profile-public-published-tests',
  templateUrl: './profile-public-published-tests.component.html',
  styleUrls: ['./profile-public-published-tests.component.css']
})
export class ProfilePublicPublishedTestsComponent extends ProfileComponent implements OnInit {

  constructor(session: SessionService, router: Router, route: ActivatedRoute,
    userS: UserService, data: DataService, authService: AuthService) {
      super(session, router, route, userS, data, authService)
  }

  ngOnInit(): void {
  }

}
