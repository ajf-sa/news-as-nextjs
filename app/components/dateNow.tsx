export const DateNow = () => {
    var d = new Date(),
        month = '' + (d.getMonth() + 1),
        day = '' + d.getDate(),
        year = d.getFullYear();

    if (month.length < 2)
        month = '0' + month;
    if (day.length < 2)
        day = '0' + day;
    const result = [year, month, day].join('-')

    return(
        <span className="px-1"> 
        {
        result
}
         </span>
         )
}