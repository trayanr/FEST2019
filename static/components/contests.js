new Vue({
    el: "#app",
    data: {
        show: "all",

    },
    created() {
        setTimeout(() => {
            //За диаграмаите

            let charts = [
                {
                    id: 'dayChart',
                    ls: ['Иван', 'Гошо', 'Стефания'],
                    d: [199, 139, 45],
                },
                {
                    id: 'monthChart',
                    ls: ['Гергана', 'Петър', 'Страхил'],
                    d: [1990, 1390, 450],
                },
                {
                    id: 'weekChart',
                    ls: ['Димитрина', 'Богомил', 'Хвалабог'],
                    d: [3990, 2390, 4532],
                },
                {
                    id: 'yearChart',
                    ls: ['Кристина', 'Кристиян', 'Красимир'],
                    d: [39900, 23900, 45032],
                }
            ];

            charts.forEach(function (e) {
                new Chart(document.getElementById(e.id),
                    {
                        type: 'bar',
                        data: {
                            labels: e.ls,
                            datasets: [{
                                label: '# Точки',
                                data: e.d,
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
            });


        }, 100);
    }

});