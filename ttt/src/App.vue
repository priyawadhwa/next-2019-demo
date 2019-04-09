<template>
  <div>
    <h1> Tic Tac Toe! </h1>

    <div class="board">
      <div v-on:click="selected(1)" v-bind:style="{ color: colors[1]}" id="1" class="cell"> {{ one }} </div>
      <div v-on:click="selected(2)" v-bind:style="{ color: colors[2]}" id="2" class="cell"> {{ two }} </div>
      <div v-on:click="selected(3)" v-bind:style="{ color: colors[3]}" id="3" class="cell"> {{ three }} </div>

      <div v-on:click="selected(4)" v-bind:style="{ color: colors[4]}" id="4" class="cell"> {{ four }}</div>
      <div v-on:click="selected(5)" v-bind:style="{ color: colors[5]}" id="5" class="cell"> {{ five }}</div>
      <div v-on:click="selected(6)" v-bind:style="{ color: colors[6]}" id="6" class="cell"> {{ six }}</div>

      <div v-on:click="selected(7)" v-bind:style="{ color: colors[7]}" id="7" class="cell"> {{ seven }} </div>
      <div v-on:click="selected(8)" v-bind:style="{ color: colors[8]}" id="8" class="cell"> {{ eight }} </div>
      <div v-on:click="selected(9)" v-bind:style="{ color: colors[9]}" id="9" class="cell"> {{ nine }}</div>
    </div>

    <div class="winner" v-if="winningPlayer!=''"> Winning player is {{ winningPlayer }}  </div>
  </div>
  
</template>

<script>

const axios = require("axios");

export default {
  name: "app",
  components: {
  },
  created() {
  },
  data() {
    return {
      tracking: {},
      trackingLetter: {},
      currentLetter: "O",
      currentPlayer: false,
      winningPlayer: "",
      one: "",
      two: "",
      three: "",
      four: "",
      five: "",
      six: "",
      seven: "",
      eight: "",
      nine: "",
      colors: {},
      xColor: "black",
      oColor: "black",
      activeColor: "black",
    }
  },
  destroyed() {
  },
  mounted() {
    axios
      .get('http://35.236.82.179:8888/colors')
      .then(response => {
        console.log(response.data)
        this.oColor = response.data.colors.o
        this.xColor = response.data.colors.x
      }).catch(error => {
      console.log(error)
      })
    this.activeColor = this.oColor
  },
  methods: {
    selected: function(cell) {
      if (this.winningPlayer != "") {
        return
      }
      if (cell in this.tracking) {
        return
      }

      if (cell == 1) {
        this.one = this.currentLetter
      }
      if (cell == 2) {
        this.two = this.currentLetter
      }
      if (cell == 3) {
        this.three = this.currentLetter
      }
      if (cell == 4) {
        this.four = this.currentLetter
      }
      if (cell == 5) {
        this.five = this.currentLetter
      }
      if (cell == 6) {
        this.six = this.currentLetter
      }
      if (cell == 7) {
        this.seven = this.currentLetter
      }
      if (cell == 8) {
        this.eight = this.currentLetter
      }
      if (cell == 9) {
        this.nine = this.currentLetter
      }

      this.tracking[cell] = this.currentPlayer
      this.currentPlayer = !this.currentPlayer
      this.trackingLetter[cell] = this.currentLetter
      if (this.currentPlayer) {
        this.currentLetter = "X"
        this.colors[cell] = this.xColor
        this.activeColor = this.xColor
      } else {
        this.currentLetter = "O"
        this.colors[cell] = this.oColor
        this.activeColor = this.oColor
      }

      let winner = this.winner()
      if (winner != "") {
        this.winningPlayer = winner
        console.log("winner is", winner)
      }
    },
    winner: function() {
      if(1 in this.tracking && 2 in this.tracking && 3 in this.tracking) {
        if (this.tracking[1] && this.tracking[2] && this.tracking[3]) {
          return "X"
        }
        if (!this.tracking[1] && !this.tracking[2] && !this.tracking[3]) {
          return "O"
        }
      }

      if(4 in this.tracking && 5 in this.tracking && 6 in this.tracking) {
        if (this.tracking[4] && this.tracking[5] && this.tracking[6]) {
          return "X"
        }
        if (!this.tracking[4] && !this.tracking[5] && !this.tracking[6]) {
          return "O"
        }
      }

      if(7 in this.tracking && 8 in this.tracking && 9 in this.tracking) {
        if (this.tracking[7] && this.tracking[8] && this.tracking[9]) {
          return "X"
        }
        if (!this.tracking[7] && !this.tracking[8] && !this.tracking[9]) {
          return "O"
        }
      }

      if(1 in this.tracking && 4 in this.tracking && 7 in this.tracking) {
        if (this.tracking[1] && this.tracking[4] && this.tracking[7]) {
          return "X"
        }
        if (!this.tracking[1] && !this.tracking[4] && !this.tracking[7]) {
          return "O"
        }
      }

      if(2 in this.tracking && 5 in this.tracking && 8 in this.tracking) {
        if (this.tracking[2] && this.tracking[5] && this.tracking[8]) {
          return "X"
        }
        if (!this.tracking[2] && !this.tracking[5] && !this.tracking[8]) {
          return "O"
        }
      }

      if(3 in this.tracking && 6 in this.tracking && 9 in this.tracking) {
        if (this.tracking[3] && this.tracking[6] && this.tracking[9]) {
          return "X"
        }
        if (!this.tracking[3] && !this.tracking[6] && !this.tracking[9]) {
          return "O"
        }
      }

      if(1 in this.tracking && 5 in this.tracking && 9 in this.tracking) {
        if (this.tracking[1] && this.tracking[5] && this.tracking[9]) {
          return "X"
        }
        if (!this.tracking[1] && !this.tracking[5] && !this.tracking[9]) {
          return "O"
        }
      }

      if(3 in this.tracking && 5 in this.tracking && 7 in this.tracking) {
        if (this.tracking[3] && this.tracking[5] && this.tracking[7]) {
          return "X"
        }
        if (!this.tracking[3] && !this.tracking[5] && !this.tracking[7]) {
          return "O"
        }
      }

      return ""
    }
  }
};
</script>

<style>
* {
  box-sizing: border-box;
}
body {
  margin: 0;
  padding: 0;
  font-family: "Google Sans", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

h1 {
    text-align: center;
}

.board {
    border: 2px solid #333333;
    margin: auto;
    width: 60%;
    height: 60%;
    font-size: 24px;
    background-color: #EEEEEE;
}

.cell {
    position: relative;
    float: left;
    width: 33%;
    height: 200px;
    border: 2px solid #333333;
    padding-top: 12px;
    font-weight: bold;
    cursor: pointer;
    text-align: center;
    font-size: 100px;
}

.winner {
  position: absolute;
  font-size: 40px;
  font-weight: bold;
  bottom: 2%;
  left: 40%;
  height: 100px;
}

</style>
