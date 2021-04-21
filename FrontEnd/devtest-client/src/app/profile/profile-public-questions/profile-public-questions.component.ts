import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, Question, UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ProfileComponent } from '../profile.component';
@Component({
  selector: 'app-profile-public-questions',
  templateUrl: './profile-public-questions.component.html',
  styleUrls: ['./profile-public-questions.component.css']
})
export class ProfilePublicQuestionsComponent extends ProfileComponent implements OnInit {

  constructor(session: SessionService, router: Router, route: ActivatedRoute,
    userS: UserService, data: DataService, authService: AuthService) {
      super(session, router, route, userS, data, authService)
  }

  doInheritHasUserAction(){
    super.doInheritHasUserAction()
    if(this.getSessionUser().isStudent()){
        this.router.navigate(['/'])
    }
  }
  
  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.ngOnDestroy()
  }

}
