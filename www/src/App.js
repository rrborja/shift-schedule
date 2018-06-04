import React, { Component } from 'react';
import TimeTable from './time-table';
import DatePicker from 'react-datepicker'
import moment from 'moment';

import './App.css'
import "react-datepicker/dist/react-datepicker.css";

class App extends Component {
    constructor() {
        super();
        this.state = {
            startDate: moment(),
            shifts: {},
            apiDate: ""
        };
        this.handleChange = this.handleChange.bind(this);
        this.reupdate = this.reupdate.bind(this);
    }

    render() {
        return (
            <div>
                <DatePicker
                    className="date-picker-size"
                    selected={this.state.startDate}
                    onChange={this.handleChange}
                />
                <TimeTable reupdate={this.reupdate} apiDate={this.state.apiDate} shifts={this.state.shifts}/>
            </div>
        );
    }

    componentWillMount() {
        this.handleChange(this.state.startDate);
    }

    reupdate() {
        this.handleChange(this.state.startDate);
    }

    handleChange(pickedDate) {
        let date = new Date(pickedDate);

        let month = date.getMonth() + 1;
        let day = date.getDate();
        let year = date.getFullYear();

        fetch(`http://localhost:8080/`+month+'/'+day+'/'+year)
            .then(result=>result.json())
            .then(items=>{
                let arr = {};
                for (let i = 0; i < items.length; i++) {
                    arr[items[i].start] = items[i];
                }

                this.setState({startDate: pickedDate, apiDate: month+'/'+day+'/'+year,  shifts: arr});
            })
    }
}

export default App;
