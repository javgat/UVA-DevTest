import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, Question, UserService } from '@javgat/devtest-api';
import { tipoPrint } from 'src/app/shared/app.model';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { ProfileComponent } from '../profile.component';
@Component({
  selector: 'app-profile-public-questions',
  templateUrl: './profile-public-questions.component.html',
  styleUrls: ['./profile-public-questions.component.css']
})
export class ProfilePublicQuestionsComponent extends ProfileComponent implements OnInit {

  questions: Question[]
  constructor(session: SessionService, router: Router, route: ActivatedRoute,
    userS: UserService, data: DataService, authService: AuthService) {
      super(session, router, route, userS, data, authService)
      this.questions = []
  }

  doProfileAction(): void{
    this.getPublicQuestions(true)
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

  getPublicQuestions(primera: boolean){
    this.userS.getPublicEditQuestionsOfUser(this.id).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas editables públicas de un usuario", primera, this.getPublicQuestions, this)
    )
  }

  tipoPrint(tipo: string, eleccionUnica: boolean | undefined){
    tipoPrint(tipo, eleccionUnica)
  }

}
