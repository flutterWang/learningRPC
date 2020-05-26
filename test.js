import React, {Fragment,Component  } from 'react';
require('./App.css');
const numbers = [1, 2, 3, 4, 5];

class todoList extends Component {
  constructor(props) {
    super(props);
    this.state = {
        inputValue: '',
    } 
  }

    NumberList(props) {
        const numbers = props.numbers;
        const listItems = (numbers || []).map((number) =>
            <li key={number.toString()}>
            {number}
            </li>
        );
        return (
            <ul>{listItems}</ul>
        );
    }
  
   render() {
    return (
      <div class="dd">
        <div class="nav">
         <div class="box">
          <div class="style2">
            <svg className="icon" aria-hidden="true">
            <use xlinkHref="#icon-xingkongmianmao"></use>
            </svg>
            <a href="#"> 爱情 </a >
          </div>
            <div class="style1">
            <input 
              type="text" 
              value="某猪是大坏蛋"
            />
            </div>
          </div>
          <div class="style3">
              <div class="one">
                <svg className="icon" aria-hidden="true">
                <use xlinkHref="#icon-bojuelvpai-baipai"></use>
                </svg>
                <a href="#">恋爱小册</a >
              </div>
                <div class="two">
                  <svg className="icon" aria-hidden="true">
                  <use xlinkHref="#icon-bojuelvpai-baipai2"></use>
                  </svg>
                  <a href="#">下载恋爱 App</a >
                </div>
                  <div class="fin">
                   <select class="ss">
                    <option value="相识">相识</option>
                    <option value="相知">相知</option>
                    <option value="相爱">相爱</option>
                   </select>
                  </div>
           </div>
          </div>
      
         <div class="cc">
          <NumberList number={numbers} ></NumberList>
         </div>
         <div class="ee">
          <input value="submit"></input>
         </div>
         <div class="hh">
          <input value="submit"></input>
         </div>
        </div>
    )
    }
}