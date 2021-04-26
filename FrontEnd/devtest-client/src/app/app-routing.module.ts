import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminTeamsComponent } from './admin/admin-teams/admin-teams.component';
import { AdminUsersComponent } from './admin/admin-users/admin-users.component';
import { AdminComponent } from './admin/admin.component';
import { CreateTeamComponent } from './create-team/create-team.component';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { LoginComponent } from './login/login.component';
import { MainComponent } from './main/main.component';
import { ProfilePublicQuestionsComponent } from './profile/profile-public-questions/profile-public-questions.component';
import { ProfilePublicTestsComponent } from './profile/profile-public-tests/profile-public-tests.component';
import { ProfileTeamsComponent } from './profile/profile-teams/profile-teams.component';
import { ProfileComponent } from './profile/profile.component';
import { PtestComponent } from './ptest/ptest.component';
import { QuestionCreateComponent } from './question-create/question-create.component';
import { QuestionTeamsComponent } from './question-teams/question-teams.component';
import { QuestionComponent } from './question/question.component';
import { QuestionsFavUserComponent } from './questions-fav-user/questions-fav-user.component';
import { QuestionsSharedUserComponent } from './questions-shared-user/questions-shared-user.component';
import { QuestionsUserComponent } from './questions-user/questions-user.component';
import { QuestionsComponent } from './questions/questions.component';
import { RecoverPasswordComponent } from './recover-password/recover-password.component';
import { SigninComponent } from './signin/signin.component';
import { TeamComponent } from './team/team.component';
import { TeamsComponent } from './teams/teams.component';
import { TestCreateComponent } from './test-create/test-create.component';
import { TestPTestsComponent } from './test-ptests/test-ptests.component';
import { TestTeamsComponent } from './test-teams/test-teams.component';
import { TestComponent } from './test/test.component';
import { TestsFavUserComponent } from './tests-fav-user/tests-fav-user.component';
import { TestsSharedUserComponent } from './tests-shared-user/tests-shared-user.component';
import { TestsUserComponent } from './tests-user/tests-user.component';
import { TestsComponent } from './tests/tests.component';

const routes: Routes = [
  {path: 'signin', component: SigninComponent},
  {path: 'login', component: LoginComponent},
  {path: 'profile/:id', component: ProfileComponent}, //Profile: Lo ven todos
  {path: 'profile/:id/teams', component: ProfileTeamsComponent},
  {path: 'profile/:id/et', component: ProfilePublicTestsComponent},
  {path: 'profile/:id/q', component: ProfilePublicQuestionsComponent},
  {path: 'u/:username/q', component: QuestionsUserComponent}, // U (User): Lo ven el usuario y el admin
  {path: 'u/:username/sq', component: QuestionsSharedUserComponent},
  {path: 'u/:username/fq', component: QuestionsFavUserComponent},
  {path: 'u/:username/et', component: TestsUserComponent},
  {path: 'u/:username/set', component: TestsSharedUserComponent},
  {path: 'u/:username/fet', component: TestsFavUserComponent},
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
  {path: 'et', component: TestsComponent},
  {path: 'et/:testid', component: TestComponent},
  {path: 'etCreate', component: TestCreateComponent},
  {path: 'et/:testid/teams', component: TestTeamsComponent},
  {path: 'et/:testid/ptests', component: TestPTestsComponent},
  {path: 'pt/:testid', component: PtestComponent},
  {path: 'pt/:testid/q/:id', component: QuestionComponent},
  {path: 'forgotPassword', component: ForgotPasswordComponent},
  {path: 'recoverPassword/:username', component: RecoverPasswordComponent},
  {path: '**', component: MainComponent, pathMatch:'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
