import { User, Team, TeamRole } from "@javgat/devtest-api";

export enum Tipo{
    SUCCESS = "success",
    ERROR = "danger",
    WARNING= "warning",
    INFO="info"
}

export class Mensaje{
    readonly texto: String;
    readonly type : Tipo;
    readonly mostrar : boolean;
    constructor(mensaje?:String, type?:Tipo, mostrar?:boolean){
        this.texto = mensaje || "";
        this.type = type || Tipo.INFO;
        this.mostrar = mostrar || false;
    }
}

export class SessionLogin{
    readonly logged: boolean
    readonly userid: String // identificador para iniciar sesion
    constructor(logged:boolean, userid?:String){
        this.logged = logged
        this.userid = userid || ""
    }
}

export class SessionUser implements User{
    readonly isEmpty: boolean
    readonly username: string;
    readonly email: string;
    readonly fullname?: string | undefined;
    readonly type?: User.TypeEnum | undefined;
    /**
     * 
     * @param username If undefined, SessionUser isEmpty
     * @param email If undefined, SessionUser isEmpty
     * @param fullname 
     * @param type 
     */
    constructor(username?:string, email?:string, fullname?:string, type?:User.TypeEnum){
        this.username = username || ""
        this.email = email || ""
        this.fullname = fullname
        this.type = type
        this.isEmpty = (this.username == "" || this.email == "")
    }
}

export class CTeam implements Team{
    teamname: string;
    description?: string | undefined;
    constructor(teamname: string, description?: string){
        this.teamname = teamname
        this.description = description
    }
}

export class CTeamRole implements TeamRole{
    role: TeamRole.RoleEnum;
    user: SessionUser;
    constructor(role: TeamRole.RoleEnum, user: SessionUser){
        this.role = role
        this.user = user
    }
}