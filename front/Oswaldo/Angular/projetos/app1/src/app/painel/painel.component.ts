import { Component, OnInit, EventEmitter, Output, OnDestroy } from '@angular/core';
import { Frase } from '../shared/frase.model';
import { FRASES } from './frases-mock'

@Component({
  selector: 'app-painel',
  templateUrl: './painel.component.html',
  styleUrls: ['./painel.component.css']
})
export class PainelComponent implements OnInit {

  public frase: Frase[] = FRASES
  public instrucao: string = 'Traduza a frase:'
  public resposta: string = ''

  public rodada: number = 0
  public rodadaFrase!: Frase 

  public progresso: number = 0

  public tentativas: number = 3

  @Output() public encerrarJogo: EventEmitter<string> = new EventEmitter

  constructor() { 
    this.atualizaRodada()
  }

  ngOnInit(): void {
  }

  ngDestroy(): void{
  }

  atualizarResposta(resposta: Event): void{
    this.resposta = (<HTMLInputElement>resposta.target).value 
  }

  public verificarResposta(): void{

    if(this.rodadaFrase.frasePtbr == this.resposta){
      
      //trocar pergunta da rodada
      this.rodada++
      
      //progresso incrementado
      this.progresso = this.progresso + (100 / this.frase.length)

      //
      if(this.rodada === 4){
        this.encerrarJogo.emit('vitória')
      }

      //atualiza o obj rodadaFrase
      this.atualizaRodada();
  

    }else{
      //decrementar a variavel tentativas
      this.tentativas--
      if(this.tentativas === -1){
        this.encerrarJogo.emit('derrota')
      }
    }    
  }

  public atualizaRodada(): void{
    //define a frase da rodada com base em alguma lógica
    this.rodadaFrase = this.frase[this.rodada];
    
    //limpar a resposta anterior fornecida pelo usuário
    this.resposta = '';
  }



}
