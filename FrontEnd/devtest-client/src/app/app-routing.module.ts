import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminTeamsComponent } from './admin/admin-teams/admin-teams.component';
import { AdminUsersComponent } from './admin/admin-users/admin-users.component';
import { AdminComponent } from './admin/admin.component';
import { CreateTeamComponent } from './create-team/create-team.component';
import { LoginComponent } from './login/login.component';
import { MainComponent } from './main/main.component';
import { ProfileComponent } from './profile/profile.component';
import { QuestionCreateComponent } from './question-create/question-create.component';
import { QuestionTeamsComponent } from './question-teams/question-teams.component';
import { QuestionComponent } from './question/question.component';
import { QuestionsComponent } from './questions/questions.component';
import { SigninComponent } from './signin/signin.component';
import { TeamComponent } from './team/team.component';
import { TeamsComponent } from './teams/teams.component';
import { TestsComponent } from './tests/tests.component';

const routes: Routes = [
  {path: 'signin', component: SigninComponent},
  {path: 'login', component: LoginComponent},
  {path: 'profile/:id', component: ProfileComponent},
  {path: 'admin', component: AdminComponent},
  {path: 'admin/users', component: AdminUsersComponent},
  {path: 'admin/teams', component: AdminTeamsComponent},
  {path: 'ptests', component: TestsComponent},
  {path: 'teams', component: TeamsComponent},
  {path: 'teams/:id', component: TeamComponent},
  {path: 'createTeam', component: CreateTeamComponent},
  {path: 'q', component: QuestionsComponent},
  {path: 'q/:id', component: QuestionComponent},
  {path: 'q/:id/teams', component: QuestionTeamsComponent},
  {path: 'qCreate', component: QuestionCreateComponent},
  {path: '**', component: MainComponent, pathMatch:'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
