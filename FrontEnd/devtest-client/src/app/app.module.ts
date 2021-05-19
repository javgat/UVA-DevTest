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
import { LqAddQuestionTestComponent } from './list-questions/lq-add-question-test/lq-add-question-test.component';
import { LqFavsUserComponent } from './list-questions/lq-favs-user/lq-favs-user.component';
import { LtFavsUserComponent } from './list-tests/lt-favs-user/lt-favs-user.component';
import { TestsFavUserComponent } from './tests-fav-user/tests-fav-user.component';
import { QuestionsFavUserComponent } from './questions-fav-user/questions-fav-user.component';
import { RecoverPasswordComponent } from './recover-password/recover-password.component';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { PtestComponent } from './ptest/ptest.component';
import { TestPTestsComponent } from './test-ptests/test-ptests.component';
import { PtestInvitesComponent } from './ptest-invites/ptest-invites.component';
import { PtestsUserComponent } from './ptests-user/ptests-user.component';
import { PtestsSharedUserComponent } from './ptests-shared-user/ptests-shared-user.component';
import { LptUserComponent } from './list-tests/lpt-user/lpt-user.component';
import { ListPtestsComponent } from './list-tests/list-ptests/list-ptests.component';
import { LptProfilePublicComponent } from './list-tests/lpt-profile-public/lpt-profile-public.component';
import { LptSharedUserComponent } from './list-tests/lpt-shared-user/lpt-shared-user.component';
import { LptInvitedUserComponent } from './list-tests/lpt-invited-user/lpt-invited-user.component';
import { ProfilePublicPublishedTestsComponent } from './profile/profile-public-published-tests/profile-public-published-tests.component';
import { PtestsComponent } from './ptests/ptests.component';
import { PtestsInvitedUserComponent } from './ptests-invited-user/ptests-invited-user.component';
import { LptSolvableComponent } from './list-tests/lpt-solvable/lpt-solvable.component';
import { PtestsSolvableUserComponent } from './ptests-solvable-user/ptests-solvable-user.component';
import { AnsweringListPQuestionsComponent } from './answering-list-pquestions/answering-list-pquestions.component';
import { AnsweringPQuestionComponent } from './answering-pquestion/answering-pquestion.component';
import { PtestAnswersComponent } from './ptest-answers/ptest-answers.component';
import { PquestionQAnswersComponent } from './pquestion-qanswers/pquestion-qanswers.component';
import { AnswerComponent } from './answer/answer.component';
import { QanswerComponent } from './qanswer/qanswer.component';
import { ListAnswersComponent } from './list-answers/list-answers.component';
import { LaCorrectedComponent } from './list-answers/la-corrected/la-corrected.component';
import { LaUncorrectedComponent } from './list-answers/la-uncorrected/la-uncorrected.component';
import { ListQAnswersComponent } from './list-qanswers/list-qanswers.component';
import { LqaAnswerComponent } from './list-qanswers/lqa-answer/lqa-answer.component';
import { LqaQuestionComponent } from './list-qanswers/lqa-question/lqa-question.component';
import { PtestUAnswersComponent } from './ptest-u-answers/ptest-u-answers.component';
import { LaUserTestComponent } from './list-answers/la-user-test/la-user-test.component';
import { NavBarLoggedComponent } from './nav-bar/nav-bar-logged/nav-bar-logged.component';
import { NavBarNotLoggedComponent } from './nav-bar/nav-bar-not-logged/nav-bar-not-logged.component';
import { TeamResourcesComponent } from './team-resources/team-resources.component';
import { LtAdminTeamComponent } from './list-tests/lt-admin-team/lt-admin-team.component';
import { LptInvitedTeamComponent } from './list-tests/lpt-invited-team/lpt-invited-team.component';
import { LqAdminTeamComponent } from './list-questions/lq-admin-team/lq-admin-team.component';
import { AdminEmailComponent } from './admin/admin-email/admin-email.component';
import { LgiStudentComponent } from './main/logged-in/lgi-student/lgi-student.component';
import { LgiTeacherComponent } from './main/logged-in/lgi-teacher/lgi-teacher.component';
import { NvStudentComponent } from './nav-bar/nav-bar-logged/nv-student/nv-student.component';
import { NvTeacherAdminComponent } from './nav-bar/nav-bar-logged/nv-teacher-admin/nv-teacher-admin.component';
import { LptAnsweredComponent } from './list-tests/lpt-answered/lpt-answered.component';
import { LptPendingComponent } from './list-tests/lpt-pending/lpt-pending.component';
import { PtestsAnsweredUserComponent } from './ptests-answered-user/ptests-answered-user.component';
import { PtestsPendingUserComponent } from './ptests-pending-user/ptests-pending-user.component';

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
    LtProfilePublicComponent,
    LqAddQuestionTestComponent,
    LqFavsUserComponent,
    LtFavsUserComponent,
    TestsFavUserComponent,
    QuestionsFavUserComponent,
    RecoverPasswordComponent,
    ForgotPasswordComponent,
    PtestComponent,
    TestPTestsComponent,
    PtestInvitesComponent,
    PtestsUserComponent,
    PtestsSharedUserComponent,
    LptUserComponent,
    ListPtestsComponent,
    LptProfilePublicComponent,
    LptSharedUserComponent,
    LptInvitedUserComponent,
    ProfilePublicPublishedTestsComponent,
    PtestsComponent,
    PtestsInvitedUserComponent,
    LptSolvableComponent,
    PtestsSolvableUserComponent,
    AnsweringListPQuestionsComponent,
    AnsweringPQuestionComponent,
    PtestAnswersComponent,
    PquestionQAnswersComponent,
    AnswerComponent,
    QanswerComponent,
    ListAnswersComponent,
    LaCorrectedComponent,
    LaUncorrectedComponent,
    LaUserTestComponent,
    ListQAnswersComponent,
    LqaAnswerComponent,
    LqaQuestionComponent,
    PtestUAnswersComponent,
    NavBarLoggedComponent,
    NavBarNotLoggedComponent,
    TeamResourcesComponent,
    LtAdminTeamComponent,
    LptInvitedTeamComponent,
    LqAdminTeamComponent,
    AdminEmailComponent,
    LgiStudentComponent,
    LgiTeacherComponent,
    NvStudentComponent,
    NvTeacherAdminComponent,
    LptAnsweredComponent,
    LptPendingComponent,
    PtestsAnsweredUserComponent,
    PtestsPendingUserComponent
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

