import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-button',
  templateUrl: './button.component.html',
  styleUrls: ['./button.component.css']
})
export class ButtonComponent implements OnInit {

  @Input() text = "";
  @Input() color = "";
  @Output() handleBtnClick = new EventEmitter()

  constructor() { }

  ngOnInit(): void {
  }

  handleOnClick() {
    this.handleBtnClick.emit()
  }

}
