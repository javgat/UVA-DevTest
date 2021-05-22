import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TeamService, TipoRol, TiporolService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, Tipo } from 'src/app/shared/app.model';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { AdminComponent } from '../admin.component';

@Component({
  selector: 'app-admin-permissions',
  templateUrl: './admin-permissions.component.html',
  styleUrls: ['./admin-permissions.component.css']
})
export class AdminPermissionsComponent extends AdminComponent implements OnInit {

  private tipoRoles: TipoRol[]
  private tipoRolesSubscription: Subscription

  tipoNoReg?: TipoRol

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, teamS: TeamService, private trS: TiporolService) {
    super(session, router, data, userS, teamS)
    this.tipoRoles = []
    this.tipoRolesSubscription = this.session.sessionTipoRoles.subscribe(
        valor => {
            this.tipoRoles = valor
            this.setPermissionData()
        }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.tipoRolesSubscription.unsubscribe()
    super.onDestroy()
  }

  getTipoRolNoRegistrado(): TipoRol{
    let noregs = this.tipoRoles.filter(tipo => tipo.rolBase == TipoRol.RolBaseEnum.NoRegistrado)
    return noregs[0]
  }

  setPermissionData(){
    this.tipoNoReg = this.getTipoRolNoRegistrado()
  }

  enviarNoReg(){
    this.sendNoReg(true)
  }

  sendNoReg(primera: boolean){
    if(this.tipoNoReg==undefined) return
    this.trS.putTipoRol(this.tipoNoReg.nombre, this.tipoNoReg).subscribe(
      resp=>{
        this.setPermissionData()
        this.cambiarMensaje(new Mensaje("Datos de permisos actualizados con éxito", Tipo.SUCCESS, true))
      },
      err => this.handleErrRelog(err, "cambiar permisos tipo no registrado", primera, this.sendNoReg, this)
    )
  }

}
