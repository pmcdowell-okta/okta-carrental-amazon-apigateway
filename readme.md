![Okta Rental](images/oktarental.png)

##### Renting the right car at the right price is almost as important as security.

The sample code below can be used to provide backend services for the Okta Car Rental Sample application.

The project relies on Serverless Framework, which is availble here: https://serverless.com

The Serverless framework is not required, But it makes the deployment **Much** easier

If you wish to recompile or change the code, it is written in **GO**, so you will neeed to set that up, however the binaries required to deploy this are included, so it *should* run out-of-box

####Deployment Instructions

* Clone the repository like this <br/> 
`git clone https://github.com/pmcdowell-okta/okta-carrental-amazon-apigateway.git`

* If you have GO Language installed, you can compile the code like this, if you are not compiling the code, you can **Skip** this step, don't sweat it <br/>
`make setup`   
`make build`<br/>
The steps above will pull down the required libraries, and build the code

* Deploy the code using Serverless Framework <br/>
`sls deploy`

If everything works as it should, this would push the required Amazon API Gateway, Custom Authorizer, and supporting Lambda Functions. The output should look something like this:

<code><br/>
Serverless: Packaging service...<br/>
Serverless: Excluding development dependencies...<br/>
Serverless: Uploading CloudFormation file to S3...<br/>
Serverless: Uploading artifacts...<br/>
Serverless: Uploading service .zip file to S3 (9.06 MB)...<br/>
Serverless: Validating template...<br/>
Serverless: Updating Stack...<br/>
Serverless: Checking Stack update progress...<br/>
..........................<br/>
Serverless: Stack update finished...<br/>
Service Information<br/>
service: golang-auth5<br/>
stage: dev<br/>
region: us-east-1<br/>
stack: golang-auth5-dev<br/>
api keys:<br/>
  None<br/>
endpoints:<br/>
  GET - https://pn5vw8stw7.execute-api.us-east-1.amazonaws.com/dev/vehicles<br/>
  POST - https://pn5vw8stw7.execute-api.us-east-1.amazonaws.com/dev/bookings<br/>
functions:<br/>
  authorizer: golang-auth5-dev-authorizer<br/>
  vehicles: golang-auth5-dev-vehicles<br/>
  bookings: golang-auth5-dev-bookings<br/>
</code>



