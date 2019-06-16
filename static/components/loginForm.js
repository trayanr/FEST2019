new Vue({
    el: "#app",
    data: function () {
        return {
            username: '',
            password: '',
            showError: false,
        }
    },
    created() {
        setTimeout(() => {
            //За диаграмата
            new Chart(document.getElementById('myChart'),
                {
                    type: 'bar',
                    data: {
                        labels: ['Иван', 'Гошо', 'Стефания'],
                        datasets: [{
                            label: '# Точки',
                            data: [19, 13, 4],
                            backgroundColor: [
                                'rgba(255, 99, 132, 0.2)',
                                'rgba(54, 162, 235, 0.2)',
                                'rgba(255, 206, 86, 0.2)',
                            ],
                            borderColor: [
                                'rgba(255, 99, 132, 1)',
                                'rgba(54, 162, 235, 1)',
                                'rgba(255, 206, 86, 1)',
                            ],
                            borderWidth: 1
                        }]
                    },
                    options: {
                        scales: {
                            yAxes: [{
                                ticks: {
                                    beginAtZero: true
                                }
                            }]
                        }
                    }
                });

            //За count down-a
            $('.countdown').FlipClock({

            });

        },100);

    },
    methods: {
        login(){
            var username = this.username
            var password = this.password
            var vue = this
            var d = JSON.stringify({
                username: username,
                password: password,
            })
            axios.post('/api/login', d)
            .then(function (response) {
                window.location.href = "/home"
            })
            .catch(function (error) {
                vue.showError = true
            })
        }
    }
});