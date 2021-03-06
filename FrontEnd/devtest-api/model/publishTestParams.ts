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


export interface PublishTestParams { 
    title: string;
    /**
     * True si no necesitaras invitacion para hacer el test
     */
    accesoPublico: boolean;
    autoCorrect: boolean;
    visibilidad: PublishTestParams.VisibilidadEnum;
    maxMinutes: number;
    tiempoEstricto: boolean;
}
export namespace PublishTestParams {
    export type VisibilidadEnum = 'alEntregar' | 'alCorregir' | 'manual';
    export const VisibilidadEnum = {
        AlEntregar: 'alEntregar' as VisibilidadEnum,
        AlCorregir: 'alCorregir' as VisibilidadEnum,
        Manual: 'manual' as VisibilidadEnum
    };
}
