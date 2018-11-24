import React, { Component } from "react";
import { Line } from "react-chartjs-2";
import moment from "moment";

export default class HistoryChart extends Component {
  render() {
    const { selectedId, data } = this.props;

    const labels = data[selectedId].map(e => {
      return moment.unix(e.date).format("hh:mm:ss");
    });

    const values = data[selectedId].map(e => {
      return parseFloat(e.value);
    });

    const chartData = {
      labels: labels,
      datasets: [
        {
          data: values
        }
      ]
    };

    const option = {
      title: {
        display: true,
        text: "asd"
      }
    };

    return <Line data={chartData} option={option} height={100} />;
  }
}
