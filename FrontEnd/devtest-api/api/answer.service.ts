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
/* tslint:disable:no-unused-variable member-ordering */

import { Inject, Injectable, Optional }                      from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams,
         HttpResponse, HttpEvent }                           from '@angular/common/http';
import { CustomHttpUrlEncodingCodec }                        from '../encoder';

import { Observable }                                        from 'rxjs';

import { Answer } from '../model/answer';
import { Question } from '../model/question';
import { QuestionAnswer } from '../model/questionAnswer';
import { Review } from '../model/review';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class AnswerService {

    protected basePath = 'https://localhost/DevTest';
    public defaultHeaders = new HttpHeaders();
    public configuration = new Configuration();

    constructor(protected httpClient: HttpClient, @Optional()@Inject(BASE_PATH) basePath: string, @Optional() configuration: Configuration) {
        if (basePath) {
            this.basePath = basePath;
        }
        if (configuration) {
            this.configuration = configuration;
            this.basePath = basePath || configuration.basePath || this.basePath;
        }
    }

    /**
     * @param consumes string[] mime-types
     * @return true: consumes contains 'multipart/form-data', false: otherwise
     */
    private canConsumeForm(consumes: string[]): boolean {
        const form = 'multipart/form-data';
        for (const consume of consumes) {
            if (form === consume) {
                return true;
            }
        }
        return false;
    }


    /**
     * Finishes an answer
     * Finishes an answers
     * @param answerid Id of the answer
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public finishAnswer(answerid: number, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public finishAnswer(answerid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public finishAnswer(answerid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public finishAnswer(answerid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling finishAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.put<any>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}`,
            null,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns an answers
     * Returns an answers
     * @param answerid Id of the answer
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getAnswer(answerid: number, observe?: 'body', reportProgress?: boolean): Observable<Answer>;
    public getAnswer(answerid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Answer>>;
    public getAnswer(answerid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Answer>>;
    public getAnswer(answerid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling getAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<Answer>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns all answers
     * Returns all answers
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getAnswers(observe?: 'body', reportProgress?: boolean): Observable<Array<Answer>>;
    public getAnswers(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Answer>>>;
    public getAnswers(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Answer>>>;
    public getAnswers(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<Array<Answer>>(`${this.basePath}/answers`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns an answer&#39;s questionAnswer
     * Returns an answers&#39;s questionAnswer
     * @param answerid Id of the answer
     * @param questionid Id of the question it is answering
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getQuestionAnswerFromAnswer(answerid: number, questionid: number, observe?: 'body', reportProgress?: boolean): Observable<QuestionAnswer>;
    public getQuestionAnswerFromAnswer(answerid: number, questionid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<QuestionAnswer>>;
    public getQuestionAnswerFromAnswer(answerid: number, questionid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<QuestionAnswer>>;
    public getQuestionAnswerFromAnswer(answerid: number, questionid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling getQuestionAnswerFromAnswer.');
        }

        if (questionid === null || questionid === undefined) {
            throw new Error('Required parameter questionid was null or undefined when calling getQuestionAnswerFromAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<QuestionAnswer>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/qanswers/${encodeURIComponent(String(questionid))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns an answer&#39;s questionAnswers
     * Returns an answers&#39;s questionAnswers
     * @param answerid Id of the answer
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getQuestionAnswersFromAnswer(answerid: number, observe?: 'body', reportProgress?: boolean): Observable<Array<QuestionAnswer>>;
    public getQuestionAnswersFromAnswer(answerid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<QuestionAnswer>>>;
    public getQuestionAnswersFromAnswer(answerid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<QuestionAnswer>>>;
    public getQuestionAnswersFromAnswer(answerid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling getQuestionAnswersFromAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<Array<QuestionAnswer>>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/qanswers`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns an answer question&#39;s questionAnswers. It must be only one
     * Returns an answer question&#39;s questionAnswers. It must be only one
     * @param answerid Id of the answer
     * @param questionid Id of the question
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getQuestionAnswersFromAnswerAndQuestion(answerid: number, questionid: number, observe?: 'body', reportProgress?: boolean): Observable<Array<QuestionAnswer>>;
    public getQuestionAnswersFromAnswerAndQuestion(answerid: number, questionid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<QuestionAnswer>>>;
    public getQuestionAnswersFromAnswerAndQuestion(answerid: number, questionid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<QuestionAnswer>>>;
    public getQuestionAnswersFromAnswerAndQuestion(answerid: number, questionid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling getQuestionAnswersFromAnswerAndQuestion.');
        }

        if (questionid === null || questionid === undefined) {
            throw new Error('Required parameter questionid was null or undefined when calling getQuestionAnswersFromAnswerAndQuestion.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<Array<QuestionAnswer>>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/questions/${encodeURIComponent(String(questionid))}/qanswers`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns all published questions in the test related to the answer. The DTOs will contain isRespondida
     * Returns all published questions in the test related to the answer. The DTOs will contain isRespondida
     * @param answerid Id of the answer
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getQuestionsFromAnswer(answerid: number, observe?: 'body', reportProgress?: boolean): Observable<Array<Question>>;
    public getQuestionsFromAnswer(answerid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Question>>>;
    public getQuestionsFromAnswer(answerid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Question>>>;
    public getQuestionsFromAnswer(answerid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling getQuestionsFromAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<Array<Question>>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/questions`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Answers a question
     * Answers a question
     * @param answerid Id of the answer
     * @param questionAnswer QuestionAnswer to post
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public postQuestionAnswer(answerid: number, questionAnswer: QuestionAnswer, observe?: 'body', reportProgress?: boolean): Observable<QuestionAnswer>;
    public postQuestionAnswer(answerid: number, questionAnswer: QuestionAnswer, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<QuestionAnswer>>;
    public postQuestionAnswer(answerid: number, questionAnswer: QuestionAnswer, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<QuestionAnswer>>;
    public postQuestionAnswer(answerid: number, questionAnswer: QuestionAnswer, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling postQuestionAnswer.');
        }

        if (questionAnswer === null || questionAnswer === undefined) {
            throw new Error('Required parameter questionAnswer was null or undefined when calling postQuestionAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.post<QuestionAnswer>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/qanswers`,
            questionAnswer,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Updates a QuestionAnswer
     * Updates a QuestionAnswer
     * @param answerid Id of the answer
     * @param questionid Id of the question it is answering
     * @param questionAnswer QuestionAnswer Updated
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putQuestionAnswerFromAnswer(answerid: number, questionid: number, questionAnswer: QuestionAnswer, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public putQuestionAnswerFromAnswer(answerid: number, questionid: number, questionAnswer: QuestionAnswer, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public putQuestionAnswerFromAnswer(answerid: number, questionid: number, questionAnswer: QuestionAnswer, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public putQuestionAnswerFromAnswer(answerid: number, questionid: number, questionAnswer: QuestionAnswer, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling putQuestionAnswerFromAnswer.');
        }

        if (questionid === null || questionid === undefined) {
            throw new Error('Required parameter questionid was null or undefined when calling putQuestionAnswerFromAnswer.');
        }

        if (questionAnswer === null || questionAnswer === undefined) {
            throw new Error('Required parameter questionAnswer was null or undefined when calling putQuestionAnswerFromAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.put<any>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/qanswers/${encodeURIComponent(String(questionid))}`,
            questionAnswer,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Updates an answer review
     * Updates an answer review
     * @param answerid Id of the answer
     * @param questionid Id of the question it is answering
     * @param review Review Updated
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putReview(answerid: number, questionid: number, review: Review, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public putReview(answerid: number, questionid: number, review: Review, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public putReview(answerid: number, questionid: number, review: Review, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public putReview(answerid: number, questionid: number, review: Review, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (answerid === null || answerid === undefined) {
            throw new Error('Required parameter answerid was null or undefined when calling putReview.');
        }

        if (questionid === null || questionid === undefined) {
            throw new Error('Required parameter questionid was null or undefined when calling putReview.');
        }

        if (review === null || review === undefined) {
            throw new Error('Required parameter review was null or undefined when calling putReview.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.put<any>(`${this.basePath}/answers/${encodeURIComponent(String(answerid))}/qanswers/${encodeURIComponent(String(questionid))}/review`,
            review,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Starts a new answer
     * Starts a new answer
     * @param username Username of the user who can answer the publishedTest
     * @param testid Id of the test
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public startAnswer(username: string, testid: number, observe?: 'body', reportProgress?: boolean): Observable<Answer>;
    public startAnswer(username: string, testid: number, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Answer>>;
    public startAnswer(username: string, testid: number, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Answer>>;
    public startAnswer(username: string, testid: number, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (username === null || username === undefined) {
            throw new Error('Required parameter username was null or undefined when calling startAnswer.');
        }

        if (testid === null || testid === undefined) {
            throw new Error('Required parameter testid was null or undefined when calling startAnswer.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.post<Answer>(`${this.basePath}/users/${encodeURIComponent(String(username))}/solvableTests/${encodeURIComponent(String(testid))}/answers`,
            null,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}
