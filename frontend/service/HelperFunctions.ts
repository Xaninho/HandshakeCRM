export default {
    toStringDateFormat(_date: string, _timeFlag: boolean): string {
        let dateString = "";

        if (_date != undefined && _date != null && _date != "") {

            const dateObj = new Date(`${_date}`); // receive in UTC format
            const offsetMs = dateObj.getTimezoneOffset() * 60 * 1000; // get local timezone offset in milliseconds
            const localDateObj = new Date(dateObj.getTime() - offsetMs); // adjust date for local timezone offset
            dateString = new Date(localDateObj).toISOString().split('T')[0];

            if (_timeFlag) {
                dateString = dateString + " " +
                    new Date(_date).toLocaleTimeString("it-IT", { hour: "2-digit", minute: "2-digit", });
            }
        }
        return dateString;
    },
}
