import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Team, TeamService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.css']
})
export class TeamComponent implements OnInit {


  private routeSub: Subscription
  id: string | undefined
  team : Team = {
    teamname: ""
  }

  constructor(private route: ActivatedRoute, private data: DataService,
    private teamService: TeamService) {
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.data.borrarMensaje()
      if (this.id != "" && this.id != undefined) {
        this.getTeam(this.id)
      }
    });
   }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.routeSub.unsubscribe()
    this.data.borrarMensaje()
  }

  getTeam(id: string){
    this.teamService.getTeam(id).subscribe(
      resp => {
        this.team = resp
        // Aqui pillar usuarios y roles
      },
      err => {
        this.data.cambiarMensaje(new Mensaje("No se pudo obtener el equipo: "+err.status, Tipo.ERROR, true))
      }
    )
  }

}
