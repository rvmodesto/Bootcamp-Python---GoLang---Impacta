import { Component, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { Coracao } from '../shared/coracao.model';

@Component({
  selector: 'app-tentativas',
  templateUrl: './tentativas.component.html',
  styleUrls: ['./tentativas.component.css']
})

export class TentativasComponent implements OnChanges, OnInit {

  public coracoes: Coracao[] = [
    new Coracao(true),
    new Coracao(true),
    new Coracao(true)
  ]

  @Input() public tentativas!: number

  constructor() {  
   }

  ngOnChanges(changes: SimpleChanges): void {
    // this.tentativas
    //this.coracoes.length
    if(this.tentativas !==this.coracoes.length){
      
      //teste de mesa
      //3 - 2 = 1 - 1 =0
      //3 - 1 = 2 - 1 = 1 - 1 =0
      //3 - 0 = 3 - 1 = 2

      let indice = this.coracoes.length - this.tentativas
      this.coracoes[indice -1].cheio = false
    }
   }

  ngOnInit(): void {}

}
