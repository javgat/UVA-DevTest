import { User, Team } from "@javgat/devtest-api";

export enum Tipo {
    SUCCESS = "success",
    ERROR = "danger",
    WARNING = "warning",
    INFO = "info"
}

export class Mensaje {
    readonly texto: String;
    readonly type: Tipo;
    readonly mostrar: boolean;
    constructor(mensaje?: String, type?: Tipo, mostrar?: boolean) {
        this.texto = mensaje || "";
        this.type = type || Tipo.INFO;
        this.mostrar = mostrar || false;
    }
}

export class SessionLogin {
    private readonly logged: boolean
    private readonly username: string
    constructor(logged: boolean, username?: string) {
        this.logged = logged
        this.username = username || ""
    }

    isLoggedIn(): boolean {
        return this.logged
    }

    getUserUsername(): string {
        return this.username
    }
}

export class SessionUser implements User {
    readonly empty: boolean
    readonly username: string;
    readonly email: string;
    readonly fullname: string;
    readonly rol: User.RolEnum;
    /**
     * 
     * @param username If undefined, SessionUser isEmpty
     * @param email If undefined, SessionUser isEmpty
     * @param fullname 
     * @param type 
     */
    constructor(username?: string, email?: string, fullname?: string, rol?: User.RolEnum) {
        this.username = username || ""
        this.email = email || ""
        this.fullname = fullname || ""
        this.rol = rol || User.RolEnum.Estudiante
        this.empty = (this.username == "" || this.email == "")
    }

    isEmpty(): boolean {
        return this.empty
    }

    getUsername(): string {
        return this.username
    }

    getEmail(): string {
        return this.email
    }

    getFullname(): string {
        return this.fullname
    }

    getRol(): User.RolEnum {
        return this.rol
    }
}

export class Usuario extends SessionUser {
    /**
     * 
     * @param username
     * @param email
     * @param fullname 
     * @param type 
     */
    constructor(username: string, email: string, fullname: string, rol: User.RolEnum) {
        super(username, email, fullname, rol)
    }
}

export class Equipo implements Team {
    teamname: string;
    description: string;
    soloProfesores: boolean;
    constructor(teamname: string, description: string, soloProfesores: boolean) {
        this.teamname = teamname
        this.description = description
        this.soloProfesores = soloProfesores
    }

    getTeamname(): string {
        return this.teamname
    }

    getDescription(): string {
        return this.description
    }

    isSoloProfesores(): boolean {
        return this.soloProfesores
    }
}