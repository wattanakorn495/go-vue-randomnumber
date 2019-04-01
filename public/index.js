var app = new Vue({
	el: '#app',
	data: {
		number: "random number"
	},
	methods: {
		random: function () {
			axios.get('random').then(response => (this.number = response.data))
		}
	}
})

app.number = "random number"

// app.seen = true;
// app.todos.push({text: 'new item'})