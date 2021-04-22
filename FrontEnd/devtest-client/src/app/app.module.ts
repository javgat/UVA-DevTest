import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SigninComponent } from './signin/signin.component';
import { LoginComponent } from './login/login.component';

import { HttpClientModule, HttpClientXsrfModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { ApiModule, BASE_PATH, Configuration, ConfigurationParameters } from '@javgat/devtest-api';
import { environment } from '../environments/environment';
import { MainComponent } from './main/main.component';
import { LoggedInComponent } from './main/logged-in/logged-in.component';
import { NotLoggedInComponent } from './main/not-logged-in/not-logged-in.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { ProfileComponent } from './profile/profile.component';
import { TestsComponent } from './tests/tests.component';
import { TeamsComponent } from './teams/teams.component';
import { AdminComponent } from './admin/admin.component';
import { TeamComponent } from './team/team.component';
import { CreateTeamComponent } from './create-team/create-team.component';
import { AdminUsersComponent } from './admin/admin-users/admin-users.component';
import { AdminTeamsComponent } from './admin/admin-teams/admin-teams.component';
import { QuestionsComponent } from './questions/questions.component';
import { QuestionComponent } from './question/question.component';
import { QuestionCreateComponent } from './question-create/question-create.component';
import { QuestionTeamsComponent } from './question-teams/question-teams.component';
import { QuestionsSharedUserComponent } from './questions-shared-user/questions-shared-user.component';
import { QuestionsUserComponent } from './questions-user/questions-user.component';
import { TestComponent } from './test/test.component';
import { TestCreateComponent } from './test-create/test-create.component';
import { TestTeamsComponent } from './test-teams/test-teams.component';
import { TestsUserComponent } from './tests-user/tests-user.component';
import { TestsSharedUserComponent } from './tests-shared-user/tests-shared-user.component';
import { CustomHttpInterceptorService } from './shared/interceptor.service';
import { ProfileTeamsComponent } from './profile/profile-teams/profile-teams.component';
import { ProfilePublicQuestionsComponent } from './profile/profile-public-questions/profile-public-questions.component';
import { ProfilePublicTestsComponent } from './profile/profile-public-tests/profile-public-tests.component';
import { ListQuestionsComponent } from './list-questions/list-questions.component';
import { LqSharedUserComponent } from './list-questions/lq-shared-user/lq-shared-user.component';
import { LqUserComponent } from './list-questions/lq-user/lq-user.component';
import { LqProfilePublicComponent } from './list-questions/lq-profile-public/lq-profile-public.component';
import { ListTestsComponent } from './list-tests/list-tests.component';
import { LtUserComponent } from './list-tests/lt-user/lt-user.component';
import { LtSharedUserComponent } from './list-tests/lt-shared-user/lt-shared-user.component';
import { LtProfilePublicComponent } from './list-tests/lt-profile-public/lt-profile-public.component';

export function apiConfigFactory (): Configuration {
  const params: ConfigurationParameters = {
    withCredentials: true,
  }
  return new Configuration(params);
}


@NgModule({
  declarations: [
    AppComponent,
    SigninComponent,
    LoginComponent,
    MainComponent,
    LoggedInComponent,
    NotLoggedInComponent,
    NavBarComponent,
    ProfileComponent,
    TestsComponent,
    TeamsComponent,
    AdminComponent,
    TeamComponent,
    CreateTeamComponent,
    AdminUsersComponent,
    AdminTeamsComponent,
    QuestionsComponent,
    QuestionComponent,
    QuestionCreateComponent,
    QuestionTeamsComponent,
    QuestionsSharedUserComponent,
    QuestionsUserComponent,
    TestComponent,
    TestCreateComponent,
    TestTeamsComponent,
    TestsUserComponent,
    TestsSharedUserComponent,
    ProfileTeamsComponent,
    ProfilePublicQuestionsComponent,
    ProfilePublicTestsComponent,
    ListQuestionsComponent,
    LqSharedUserComponent,
    LqUserComponent,
    LqProfilePublicComponent,
    ListTestsComponent,
    LtUserComponent,
    LtSharedUserComponent,
    LtProfilePublicComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ApiModule.forRoot(apiConfigFactory),
    FormsModule,
  ],
  providers: [
    { provide: BASE_PATH, useValue: environment.API_BASE_PATH },
    { provide: HTTP_INTERCEPTORS, useClass: CustomHttpInterceptorService, multi: true },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

