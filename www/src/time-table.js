import React, { Component } from 'react';
import './time-table.css';
import ShiftButton from './shift-button';

class TimeTable extends Component {
    constructor(props) {
        super(props);
        this.state = {
            apiDate: props.apiDate,
            shifts: this.props.shifts
        };
    }

    render() {
        let items = this.props.shifts
        var rows = []
        var next = null
        for (var i=0; i<48; i++) {
            var minutes = (i%2) * 30
            var hours = Math.floor(i/2)

            if (minutes === 0) minutes = "0" + minutes
            if (hours < 10) hours = "0" + hours

            if (items[i]) {
                next = i + items[i].end - items[i].start
            }

            rows.push(
                <tr id="time-table">
                    <td>{hours}:{minutes}</td>
                    {
                        items[i] ?
                            <td className="worker" rowSpan={items[i].end - items[i].start}>{items[i].name} ({items[i].id})</td> :
                                i>=next ? <ShiftButton reupdate={this.props.reupdate} apiDate={this.props.apiDate} timeslot={i}/> : ""
                    }
                </tr>
            )
        }
        return (
            <tbody>
            {rows}
            </tbody>
        );
    }
}

export default TimeTable;