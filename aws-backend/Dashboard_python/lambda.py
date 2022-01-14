import json
import boto3
import time
import math

from boto3.dynamodb.conditions import Key
from dashboard import dataHandler


def lambda_handler(event, context):

    agent = dataHandler.DataHandler()
    agent.get_msg_number()
    agent.get_agent_dashboard()
