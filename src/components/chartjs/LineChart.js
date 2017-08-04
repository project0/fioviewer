import Vue from 'vue'
import Chart from 'chart.js'
import merge from 'lodash/fp/merge'
import momemt from 'moment'
import Colors from './Colors'

export default Vue.extend({
  render: function (createElement) {
    return createElement(
      'canvas', {
        attrs: {
          width: '100%',
          height: '100%'
        },
        ref: 'canvas'
      }
    )
  },

  props: {
    chartData: {},
    options: {}
  },

  data () {
    return {
      defaultOptions: {
        responsive: true,
        maintainAspectRatio: false, // do not use whole page
        hover: {
          enabled: true,
          mode: 'index',
          intersect: true
        },
        tooltips: {
          enabled: true,
          mode: 'index',
          intersect: true
        },
        legend: {
          display: true
        },
        scales: {
          xAxes: [{
            ticks: {
              beginAtZero: false,
              callback: function (value, index, values) {
                return value
              }
            },
            gridLines: {
              display: true
            },
            type: 'time',
            time: {
              parser: function (v) {
                return momemt(v).utc() // remove any local timezone offsets
              },
              displayFormats: {
                'millisecond': 'SSSS [ms]',
                'second': 'mm:ss [m]',
                'minute': 'H:mm [h]',
                'hour': 'H:mm [h]'
              }
            }
          }],
          yAxes: [{
            gridLines: {
              display: true
            }
          }]
        }
      }
    }
  },

  methods: {
    renderChart (data, options) {
      let chartOptions = merge(this.defaultOptions, options)
      // ensure old chart is gone, this fixes some issue with overlaying old charts..
      if (this._chart) {
        this._chart.destroy()
      }

      this._chart = new Chart(
        this.$refs.canvas.getContext('2d'), {
          type: 'line',
          data: this.injectColor(data),
          options: chartOptions
        }
      )
      this._chart.generateLegend()
    },
    forceUpdate (newData, chart) {
      newData.datasets.forEach((dataset, i) => {
        chart.data.datasets[i].data = dataset.data
      })

      chart.data.labels = newData.labels
      chart.update()
    },
    forceRender () {
      this.renderChart(this.chartData, this.options)
    },
    injectColor (data) {
      if (!data.datasets) {
        return data
      }

      data.datasets = data.datasets.map((dataset, index) => {
        return merge(dataset, Colors[index])
      })
      return data
    }
  },
  beforeDestroy () {
    this._chart.destroy()
  },
  mounted () {
    this.renderChart(this.chartData, this.options)
  },
  watch: {
    'chartData': {
      handler (newData, oldData) {
        if (oldData) {
          let chart = this._chart

          let newDataLabels = newData.datasets.map((dataset) => {
            return dataset.label
          })

          let oldDataLabels = oldData.datasets.map((dataset) => {
            return dataset.label
          })

          if (JSON.stringify(newDataLabels) === JSON.stringify(oldDataLabels)) {
            this.forceUpdate(newData, chart)
          } else {
            this.forceRender()
          }
        }
      }
    }
  }
})
