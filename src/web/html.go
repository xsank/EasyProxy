package web

var StatisticHtml = `
<h1>EasyProxy server statistic</h1>
<h3>Service status</h3>
<table border="5">
  <thead>
    <tr>
      <th>service</th>
       <th>count</th>
      <th>status</th>
    </tr>
  </thead>
  <tbody>
	  {{range $url,$service:=.Services}}
		   <tr>
				<td>{{$service.Url}}</td>
				<td>{{$service.Count}}</td>
				<td>{{$service.Status}}</td>
		   </tr>
	  {{end}}
  </tbody>
</table>
<h3>Client status</h3>
<table border="5">
  <thead>
    <tr>
      <th>client</th>
      <th>count</th>
    </tr>
  </thead>
  <tbody>
	  {{range $host,$client :=.Clients}}
		  <tr>
			<td>{{$client.Host}}</td>
			<td>{{$client.Count}}</td>
		  </tr>
	  {{end}}
  </tbody>
</table>
`
