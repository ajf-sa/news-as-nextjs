export const DateNow = ({date}) => {
    var d = new Date(date),
        month = '' + (d.getMonth() + 1),
        day = '' + d.getDate(),
        year = d.getFullYear();

    if (month.length < 2)
        month = '0' + month;
    if (day.length < 2)
        day = '0' + day;
    const result = [day,month,year].join('-')

    return(
        <span className="px-1"> 
        {
        result
}
         </span>
         )
}