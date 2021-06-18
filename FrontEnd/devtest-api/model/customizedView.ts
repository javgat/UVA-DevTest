/**
 * DevTest
 * DevTest
 *
 * OpenAPI spec version: 1.0.0
 * Contact: javigaton@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


export interface CustomizedView { 
    rolBase: CustomizedView.RolBaseEnum;
    mensajeInicio: string;
}
export namespace CustomizedView {
    export type RolBaseEnum = 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado';
    export const RolBaseEnum = {
        Administrador: 'administrador' as RolBaseEnum,
        Profesor: 'profesor' as RolBaseEnum,
        Estudiante: 'estudiante' as RolBaseEnum,
        NoRegistrado: 'noRegistrado' as RolBaseEnum
    };
}
