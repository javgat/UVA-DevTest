import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SigninComponent } from './signin/signin.component';
import { LoginComponent } from './login/login.component';

import { HttpClientModule, HttpClientXsrfModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { ApiModule, BASE_PATH } from '@javgat/devtest-api';
import { environment } from '../environments/environment';
import { MainComponent } from './main/main.component';
import { LoggedInComponent } from './main/logged-in/logged-in.component';
import { NotLoggedInComponent } from './main/not-logged-in/not-logged-in.component';

@NgModule({
  declarations: [
    AppComponent,
    SigninComponent,
    LoginComponent,
    MainComponent,
    LoggedInComponent,
    NotLoggedInComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ApiModule,
    FormsModule,
    HttpClientXsrfModule.withOptions({
      cookieName: 'Bearer-Cookie',
      headerName: 'Bearer',
    }),
  ],
  providers: [{ provide: BASE_PATH, useValue: environment.API_BASE_PATH }],
  bootstrap: [AppComponent]
})
export class AppModule { }

