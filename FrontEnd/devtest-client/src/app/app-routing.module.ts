import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { MainComponent } from './main/main.component';
import { SigninComponent } from './signin/signin.component';

const routes: Routes = [
  {path: 'signin', component: SigninComponent},
  {path: 'login', component: LoginComponent},
  {path: '**', component: MainComponent, pathMatch:'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
