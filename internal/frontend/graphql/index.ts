import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Client = {
  __typename?: 'Client';
  avatarUrl: Scalars['String'];
  name: Scalars['String'];
};

export type ClientColors = {
  __typename?: 'ClientColors';
  client: Client;
  colors: Array<Scalars['String']>;
  error?: Maybe<Scalars['String']>;
};

export type ClientGif = {
  __typename?: 'ClientGif';
  client: Client;
  error?: Maybe<Scalars['String']>;
  gifUrl: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  clients: Array<Client>;
  level1: Array<ClientColors>;
  level2: Array<ClientColors>;
  level3: Array<ClientGif>;
};

export type _Service = {
  __typename?: '_Service';
  sdl: Scalars['String'];
};

export type ClientFragment = { __typename?: 'Client', name: string, avatarUrl: string };

export type ClientsQueryVariables = Exact<{ [key: string]: never; }>;


export type ClientsQuery = { __typename?: 'Query', clients: Array<{ __typename?: 'Client', name: string, avatarUrl: string }> };

export type Level1QueryVariables = Exact<{ [key: string]: never; }>;


export type Level1Query = { __typename?: 'Query', level1: Array<{ __typename?: 'ClientColors', colors: Array<string>, error?: string | null, client: { __typename?: 'Client', name: string, avatarUrl: string } }> };

export type Level2QueryVariables = Exact<{ [key: string]: never; }>;


export type Level2Query = { __typename?: 'Query', level2: Array<{ __typename?: 'ClientColors', colors: Array<string>, error?: string | null, client: { __typename?: 'Client', name: string, avatarUrl: string } }> };

export type Level3QueryVariables = Exact<{ [key: string]: never; }>;


export type Level3Query = { __typename?: 'Query', level3: Array<{ __typename?: 'ClientGif', gifUrl: string, error?: string | null, client: { __typename?: 'Client', name: string, avatarUrl: string } }> };

export const ClientFragmentDoc = gql`
    fragment Client on Client {
  name
  avatarUrl
}
    `;
export const ClientsDocument = gql`
    query clients {
  clients {
    ...Client
  }
}
    ${ClientFragmentDoc}`;

export function useClientsQuery(options?: Omit<Urql.UseQueryArgs<ClientsQueryVariables>, 'query'>) {
  return Urql.useQuery<ClientsQuery, ClientsQueryVariables>({ query: ClientsDocument, ...options });
};
export const Level1Document = gql`
    query level1 {
  level1 {
    client {
      ...Client
    }
    colors
    error
  }
}
    ${ClientFragmentDoc}`;

export function useLevel1Query(options?: Omit<Urql.UseQueryArgs<Level1QueryVariables>, 'query'>) {
  return Urql.useQuery<Level1Query, Level1QueryVariables>({ query: Level1Document, ...options });
};
export const Level2Document = gql`
    query level2 {
  level2 {
    client {
      ...Client
    }
    colors
    error
  }
}
    ${ClientFragmentDoc}`;

export function useLevel2Query(options?: Omit<Urql.UseQueryArgs<Level2QueryVariables>, 'query'>) {
  return Urql.useQuery<Level2Query, Level2QueryVariables>({ query: Level2Document, ...options });
};
export const Level3Document = gql`
    query level3 {
  level3 {
    client {
      ...Client
    }
    gifUrl
    error
  }
}
    ${ClientFragmentDoc}`;

export function useLevel3Query(options?: Omit<Urql.UseQueryArgs<Level3QueryVariables>, 'query'>) {
  return Urql.useQuery<Level3Query, Level3QueryVariables>({ query: Level3Document, ...options });
};