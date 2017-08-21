
import socket
import struct
import sys
import time


def main():
    """função principal do cliente ntp"""

    if len(sys.argv) < 2:
        sys.stderr.write("digite: " + sys.argv[0] + " <NTP server>")
        exit(1)
    # recebe t1, data e t4
    time1 = time.time()
    data, server= connection(sys.argv[1])
    time4 = time.time()

    #retira o stratum do vetor de bytes data
    upacket = struct.unpack('!48B', data)
    stratum = upacket[1]
    
    #retira t2 e t3 de data, corrgindo o tempo 
    time2 = struct.unpack("!12I", data)[8] + (struct.unpack("!12I", data)[9]/10000000000) - 2208988800 
    time3 = struct.unpack("!12I", data)[10] + (struct.unpack("!12I", data)[11]/10000000000) - 2208988800 

    offset = ((time2 - time1) + (time3 - time4))/2
    delay = ((time4 - time1) - (time3 - time2))/2
        
    print('Servidor: %s, stratum: %d, offset: %f, Delay: %f' % (server[0], stratum, offset, delay))

# END MAIN

def connection(server):
    """ realiza a conexão UDP e retorna os dados recebidos e o servidor de destino"""
    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    data = '\x23' + 47 * '\0'   
    data = str.encode(data)
    
    try:
        client.sendto(data, (server, 123))
        client.settimeout(1)

    except:
        print("server %s não foi encontrado (veja a conexão com a internet)" % (server))
        exit(2)

    try:
        print('mensagem enviada. Esperando resposta...')
        data, server = client.recvfrom(123)

    except socket.timeout:
        print("timed out")
        exit(3)
    
    return data, server

#END CONNECTION

if __name__ == "__main__":
    main()
