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
    TeamComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ApiModule.forRoot(apiConfigFactory),
    FormsModule,
  ],
  providers: [{provide: BASE_PATH, useValue: environment.API_BASE_PATH },],
  bootstrap: [AppComponent]
})
export class AppModule { }

