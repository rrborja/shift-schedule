import React, { Component, PropTypes } from 'react';
import ReactModal from 'react-modal';
import './shift-button.css'

const customStyles = {
    content : {
        top                   : '50%',
        left                  : '50%',
        right                 : 'auto',
        bottom                : 'auto',
        marginRight           : '-50%',
        transform             : 'translate(-50%, -50%)'
    }
};

function TimeslotString(timeslot) {
    let hours = Math.floor(timeslot/2)
    let minutes = (timeslot%2) * 30

    if (minutes === 0) minutes = "0" + minutes
    if (hours < 10) hours = "0" + hours

    return hours+":"+minutes
}

export default class ShiftButton extends Component {
    constructor (props) {
        super(props);
        this.state = {
            showModal: false,
            showErrorModal: false,
            name: "",
            id: -1,
            start: -1,
            ends: this.props.timeslot+1
        };

        this.handleOpenModal = this.handleOpenModal.bind(this);
        this.handleCloseModal = this.handleCloseModal.bind(this);
        this.handleCloseErrorModal = this.handleCloseErrorModal.bind(this);
        this.handleName = this.handleName.bind(this);
        this.handleId = this.handleId.bind(this);
        this.handleEnds = this.handleEnds.bind(this);
        this.submitShift = this.submitShift.bind(this);
    }

    handleOpenModal () {
        this.setState({ showModal: true });
    }

    handleCloseModal () {
        this.setState({ showModal: false });
    }

    handleCloseErrorModal () {
        this.setState({ showErrorModal: false });
    }

    handleName (e) {
        this.setState({ name: e.target.value });
    }

    handleId (e) {
        this.setState({ id: e.target.value });
    }

    handleEnds (e) {
        console.log(e)
        this.setState({ ends: e.target.value });
    }

    submitShift () {
        this.setState({ showModal: false });
        fetch(`http://localhost:8080/`+this.props.apiDate, {
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "name": this.state.name,
                "id": parseInt(this.state.id),
                "start": parseInt(this.props.timeslot),
                "end": parseInt(this.state.ends)
            })
        }).then(result=>{
            if (result.status === 200) {
                this.props.reupdate();
            } else {
                this.setState({ showErrorModal: true });
            }
        })
    }

    render() {
        let timeslot = this.props.timeslot;

        let availableTimeslots = [];

        for (var i=timeslot+1; i<=48; i++) {
            availableTimeslots.push(<option value={i}>{TimeslotString(i)}</option>);
        }

        return (
            <div className="main">
                <div className="vacant" onClick={this.handleOpenModal}></div>

                <ReactModal
                    isOpen={this.state.showModal}
                    contentLabel="Add Shift"
                    style={customStyles}>

                    <h2>Add Employee Shift</h2>
                    <p className="form">
                        <label for="form-name">Name</label><input id="form-name" onChange={this.handleName}/><br/>
                        <label for="form-id">ID</label><input id="form-id" onChange={this.handleId}/><br/>
                        <label for="form-starts">Selected Time</label><input id="form-starts" value={TimeslotString(timeslot)} readOnly/><br/>
                        <label for="form-ends">Shift Ends</label><select id="form-ends" onChange={this.handleEnds}>{availableTimeslots}</select><br/>
                    </p>
                    <div>
                        <div className="submit" onClick={this.submitShift}>Submit</div>
                        <div className="cancel" onClick={this.handleCloseModal}>Cancel</div>
                    </div>

                </ReactModal>

                <ReactModal
                    isOpen={this.state.showErrorModal}
                    contentLabel="Timeslots Overlapped"
                    style={customStyles}>
                    <h2>Timeslots Overlapped</h2>
                    <p>Cannot add desired timeslot because an existing shift overlaps with the selected timeslot</p>
                    <div className="submit" onClick={this.handleCloseErrorModal}>Close</div>
                </ReactModal>
            </div>
        );
    }
}