import argparse
import requests
import os
import sys
import json
import signal
import logging
import gflags
from gflags import FLAGS

def signal_handler(signum, frame):
    sys.exit(0)

def get_img_data(args):
    filenames = []
    for f in os.listdir(args.image_dir):
        if f.startswith('.'):
            continue
        filenames.append(os.path.join(args.image_dir, f))
        if len(filenames) >= args.image_num:
            break

    image_datas = []
    for fn in filenames:
        image_datas.append(open(fn, 'rb').read())
    return image_datas

def do_request(url, method, data):
    r = requests.request(method, url, data=data)
    print(r.status_code)
    if r.status_code == 200:
        return r.content

def get_img_info(args):
    image_datas = get_img_data(args)
    for data in image_datas:
        url = 'http://{}/info'.format(args.address)
        r = do_request(url, 'POST', data)
        if r:
            print(r)

if __name__ == '__main__':
    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)

    parser = argparse.ArgumentParser()
    parser.add_argument("command", help="info")
    parser.add_argument("-d", "--image_dir", help="image directory")
    parser.add_argument("-n", "--image_num", type=int,
                        help="image directory")
    parser.add_argument("-a", "--address", default="127.0.0.1:8089",
                        help="HTTP address")
    args = parser.parse_args()

    if args.command == 'info':
        get_img_info(args)
    else:
        print('illegal args')
